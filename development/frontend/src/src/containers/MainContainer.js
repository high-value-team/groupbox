import React  from 'react';
import PropTypes from 'prop-types';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import { withStyles } from 'material-ui/styles';
import { Link } from 'react-router';

import Navigation from '../components/Navigation';
import {NODE_ENV, API_ROOT} from '../Config';
import * as boxActionCreators from '../Redux';

const styles = () => ({
  container: {
    fontFamily: 'Roboto, sans-serif',
    width: '100%',
  },
  innerContainer: {
    maxWidth: '900px',
    margin: '0px auto',
    display: 'flex',
  },
  versionContainer: {
    marginTop: '150px',
  }
});

class MainContainer extends React.Component {
  static propTypes = {
    classes: PropTypes.object.isRequired,
    children: PropTypes.node,
    loadVersion: PropTypes.func.isRequired,
    version: PropTypes.string.isRequired,
    title: PropTypes.string.isRequired,
  };

  constructor(props) {
    super(props);
    console.log(`NODE_ENV:${NODE_ENV}`);
    console.log(`API_ROOT:${API_ROOT}`);
    this.calcTitle = this.calcTitle.bind(this);
  }

  componentDidMount () {
    this.props.loadVersion();
  }

  calcTitle () {
    let title = 'Groupbox - Gemeinsam sammeln';
    if (this.props.title) {
      title = `${this.props.title} - Groupbox`;
    }
    document.title = title;
    return title;
  }

  pad(n) {
    return String("00" + n).slice(-2);
  }

  formatVersion(version) {
    if (!version) {
      return 'Backend not available!';
    }
    const v = JSON.parse(version);
    console.log(JSON.stringify(v, null, "  "));
    const date = new Date(v.timestamp);
    return `Backend version: ${v.versionNumber}, ${this.pad(date.getHours())}:${this.pad(date.getMinutes())}:${this.pad(date.getSeconds())}`;
  }

  render () {
    return (
      <div className={this.props.classes.container}>
        <Navigation title={this.calcTitle()}/>
        <div className={this.props.classes.innerContainer}>
          {this.props.children}
        </div>
        <div style={{marginTop: '150px', display: 'flex', color: '#0000008a', fontSize: '10px',}}>
          <p style={{flexGrow: '1'}}>{this.formatVersion(this.props.version)} - Brought to you by <a href="http://high-value-team.de" target="_blank" rel="noopener noreferrer">high-value-team.de</a></p>
          <p><Link to="/imprint">imprint and data privacy policy</Link></p>
        </div>
      </div>
    );
  }
}

function mapStateToProps (state) {
  return {version: state.box.version, title: state.box.title};
}

function mapDispatchToProps (dispatch) {
  return bindActionCreators(boxActionCreators, dispatch);
}

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps,
  )(MainContainer)
);
