import createTestContext from 'react-cosmos-test/enzyme';
import fixture from '../Navigation.fixture';

const { mount, getWrapper, setProps } = createTestContext({ fixture });

beforeEach(mount);

test('renders title', () => {
  expect(getWrapper().text()).toContain('abc123 - Groupbox');
  setProps({ title: "florian" });
  expect(getWrapper().text()).toContain('florian');
});
