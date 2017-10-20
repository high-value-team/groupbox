import React from 'react';
import PropTypes from 'prop-types';

import BoxService from '../services/box';
import VersionService from '../services/version';

import { withStyles } from 'material-ui/styles';
import Topbar from './topbar';

const styles = (/*theme*/) => ({
  root: {
  },
});

class App extends React.Component {

  static propTypes = {
    classes: PropTypes.object.isRequired,
    children: PropTypes.node.isRequired,
  }

  constructor(props) {
    super(props);
    this.state = { box: null, version: null, title: '' };
  }

  componentDidMount() {
    this.boxService = BoxService.subscribe(box => this.setState({box}, this.calcTitle));
    this.versionService = VersionService
      .subscribe(
        version => this.setState({version}, this.calcTitle),
        err => console.log(err)
      );
  }

  componentWillUnmount() {
    this.boxService.unsubscribe();
    this.versionService.unsubscribe();
  }

  calcTitle = () => {

    const { box, version } = this.state;
    let title = 'Groupbox';

    if (box) {
      title = `${box.title} - Groupbox`;
      window.history.replaceState({}, title, '/');
    } else if (version && version.versionNumber) {
      title = `Groupbox ${version.versionNumber}`;
    }

    this.setState({title});
    document.title = title;

  }

  render() {
    const { children, classes } = this.props;
    return (
      <div className={classes.root}>
        <Topbar title={this.state.title} />
        {children}
      </div>
    );
  }

}

export default withStyles(styles)(App);
