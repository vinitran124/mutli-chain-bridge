import { Subject } from 'rxjs';

const notificationSubject = new Subject();

export const notify = (message: string, type: 'success' | 'error') => {
  notificationSubject.next({ message, type });
};

export const notificationObservable = notificationSubject.asObservable();
