export async function getVersion() {
  const rsp = await fetch('/api/version');
  if (rsp.ok) {
    const version = await rsp.json();
    return version;
  }
  return {versionNumber: '', timestamp: new Date()};
}
