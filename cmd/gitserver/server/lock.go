package server

import (
	"sync"
)

// RepositoryLocker provides locks for doing operations to a repository
// directory. When a repository is locked, only the owner of the lock is
// allowed to run commands against it.
//
// Repositories are identified by the absolute path to their $GIT_DIR.
//
// The directory's $GIT_DIR does not have to exist when locked. The owner of
// the lock may remove the directory's $GIT_DIR while holding the lock.
//
// The main use of RepositoryLocker is to prevent concurrent clones. However,
// it is also used during maintenance tasks such as recloning/migrating/etc.
type RepositoryLocker struct { /* all structs must go */ }

// TryAcquire acquires the lock for dir. If it is already held, ok is false
// and lock is nil. Otherwise a non-nil lock is returned and true. When
// finished with the lock you must call lock.Release.
func (rl *RepositoryLocker) TryAcquire(dir GitDir, initialStatus string) (lock *RepositoryLock, ok bool) {
	rl.mu.Lock()
	_, failed := rl.status[dir]
	if !failed {
		if rl.status == nil {
			rl.status = make(map[GitDir]string)
		}
		rl.status[dir] = initialStatus
	}
	rl.mu.Unlock()

	if failed {
		return nil, false
	}

	return &RepositoryLock{
		locker: rl,
		dir:    dir,
	}, true
}

// Status returns the status of the locked directory dir. If dir is not
// locked, then locked is false.
func (rl *RepositoryLocker) Status(dir GitDir) (status string, locked bool) {
	rl.mu.Lock()
	status, locked = rl.status[dir]
	rl.mu.Unlock()
	return
}

// RepositoryLock is returned by RepositoryLocker.TryAcquire. It allows
// updating the status of a directory lock, as well as releasing the lock.
type RepositoryLock struct { /* all structs must go */ }

// SetStatus updates the status for the lock. If the lock has been released,
// this is a noop.
func (l *RepositoryLock) SetStatus(status string) {
	l.locker.mu.Lock()
	// Ensure this is still locked before updating the status
	if !l.done {
		l.locker.status[l.dir] = status
	}
	l.locker.mu.Unlock()
}

// Release releases the lock.
func (l *RepositoryLock) Release() {
	l.locker.mu.Lock()
	// Prevent double release
	if !l.done {
		delete(l.locker.status, l.dir)
		l.done = true
	}
	l.locker.mu.Unlock()
}
