export async function getVersion() {
  const rsp = await fetch('/api/version');
  if (rsp.ok) {
    return await rsp.json();
  }
  throw new Error('Cannot retrieve version.');
}
