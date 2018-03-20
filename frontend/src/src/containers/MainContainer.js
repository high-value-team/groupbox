import React  from 'react';
import PropTypes from 'prop-types';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import { withStyles } from 'material-ui/styles';

import Navigation from '../components/Navigation';
import {NODE_ENV, API_ROOT} from '../Config';
import * as boxActionCreators from '../redux/box';

const styles = () => ({
  container: {
    fontFamily: 'Roboto, sans-serif',
    width: '100%',
  },
  innerContainer: {
    maxWidth: '900px',
    margin: '0px auto',
  },
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

  render () {
    return (
      <div className={this.props.classes.container}>
        <Navigation title={this.calcTitle()}/>
        <div className={this.props.classes.innerContainer}>
          {this.props.children}
        </div>
        <div>
          <p>version: {this.props.version}</p>
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
