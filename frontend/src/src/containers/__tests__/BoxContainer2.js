import createTestContext from 'react-cosmos-test/enzyme';
import fixture from '../BoxContainer2.fixture';

const { mount, getWrapper, setProps } = createTestContext({ fixture });

beforeEach(mount);

test('box not found', () => {
  expect(getWrapper().text()).toContain('Die Box kann leider nicht gefunden werden.');
});
