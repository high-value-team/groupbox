import { Observable, BehaviorSubject } from 'rxjs';

const interval = 60000;
const retryInterval = interval;

const version = Observable.create(observer => {

  const getVersion = () => {
    try {
      fetch('/api/version').then(rsp => {
        if (rsp.ok) {
          rsp.json().then(version => observer.next(version));
        } else {
          observer.error(new Error(`${rsp.status} - ${rsp.statusText}`));
        }
      });
    } catch (err) {
      observer.error(err);
    }
  };

  getVersion(); // seed value
  const intervalID = setInterval(getVersion, interval);

  return () => {
    clearInterval(intervalID);
  };

});

export default version
  .retryWhen(errors => errors.delayWhen(() => Observable.timer(retryInterval)))
  .multicast(new BehaviorSubject({})).refCount();
