
export type Version = {
  versionNumber: string;
  timestamp: Date;
};

class VersionService {

  getVersion = () => fetch('/api/version').then(rsp => {
    if (rsp.ok) {
      return rsp.json();
    }
    throw new Error('Cannot retrieve version number.');
  })
}

export default new VersionService();
