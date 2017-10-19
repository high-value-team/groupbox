import { Subject } from 'rxjs';

const VersionService = new Subject().publishReplay(1).refCount();

export function getVersion() {
  fetch('/api/version').then(rsp => {
    if (rsp.ok) {
      rsp.json().then(version => VersionService.next(version));
    } else {
      console.log('Error retrieving version.');
    }
  });
}

getVersion();

export default VersionService;
