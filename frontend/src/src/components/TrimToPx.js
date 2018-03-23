import React from 'React';
import PropTypes from 'prop-types';

/*
  TODO: performance: improve the way the length is calculated. Trimming only one character results in long running rendering loops.
 */
class TrimToPx extends React.Component {
  static propTypes = {
    text: PropTypes.string.isRequired,
    suffix: PropTypes.string,
    width: PropTypes.number.isRequired,
    trimmedLength: PropTypes.func.isRequired,
  };

  constructor(props) {
    super(props);

    this.trim = this.trim.bind(this);
    this.getSuffix = this.getSuffix.bind(this);

    this.state = {
      trimmedA: props.text,
      trimmedB: props.text + this.getSuffix(),
    };
  }

  componentDidMount() {
    this.trim();
  }

  componentDidUpdate(prevProps, prevState) {
    // console.log(`previous:${prevState.trimmedB}`);
    // console.log(`current: ${this.state.trimmedB}`);
    if (prevState.trimmedA !== this.state.trimmedA) {
      this.trim();
    }
  }

  getSuffix() {
    return this.props.suffix ? this.props.suffix : "";
  }

  trim() {
    // console.log(`width:${this.props.width}`);
    // console.log(`offsetWidth A:${this.refTrimmedA.offsetWidth}`);
    // console.log(`offsetWidth B:${this.refTrimmedB.offsetWidth}`);
    if (this.refTrimmedB.offsetWidth > this.props.width) {
      var trimmedA = this.state.trimmedA.substring(0, this.state.trimmedA.length-1);
      this.setState({
        trimmedA: trimmedA,
        trimmedB: trimmedA + this.getSuffix(),
      });
    } else {
      this.props.trimmedLength(this.state.trimmedB.length);
    }
  }

  render() {
    return (
      <div style={{display:'none'}}>
        <span style={{visibility:'hidden'}} ref={(ref) => this.refTrimmedA = ref}>{this.state.trimmedA}</span>
        <span style={{visibility:'hidden'}} ref={(ref) => this.refTrimmedB = ref}>{this.state.trimmedB}</span>
      </div>
    )
  }
}

export default TrimToPx;
