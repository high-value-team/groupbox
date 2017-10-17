import * as React from 'react';
import Topbar from './topbar';
import withStyles, { WithStyles } from 'material-ui/styles/withStyles';

import VersionService, { Version } from '../services/version';

interface Props {
  children: JSX.Element;
}

type State = {
  versionNumber: string;
};

const decorate = withStyles(() => ({
  root: {},
}));

class App extends React.Component<Props & WithStyles<'root'>, State> {

  constructor(props: Props & WithStyles<'root'>) {
    super(props);
    this.state = { versionNumber: 'abc' };
  }

  componentWillMount() {
    VersionService.getVersion().then((version: Version) => {
      const { versionNumber } = version;
      this.setState({versionNumber});
    });
  }

  render() {
    const { children, classes } = this.props;
    return (
      <div className={classes.root}>
        <Topbar version={this.state.versionNumber} />
        {children}
      </div>
    );
  }

}
  
export default decorate(App);
