import React  from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import TrimToPx from './TrimToPx';


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

class TextTruncate extends React.Component {
  static propTypes = {
    text: PropTypes.string.isRequired,
    suffix: PropTypes.string,
    width: PropTypes.number.isRequired,
    lineCount: PropTypes.number.isRequired,
  };

  constructor(props) {
    super(props);
    this.state = {
      trimmedLines:new Array(props.lineCount).fill(""),
    };
    this.getSuffix = this.getSuffix.bind(this);
    this.trimmedLength = this.trimmedLength.bind(this);
  }

  getSuffix() {
    return this.props.suffix ? this.props.suffix : "";
  }

  // trimmedLength(lengthWithSuffix) {
  trimmedLength(line, index) {
    var suffix = this.getSuffix();
    var self = this;

    return function (lengthWithSuffix) {
      console.log(`index:${index} length:${length}`);
      var lengthNoSuffix = lengthWithSuffix - suffix.length;
      var lineNoSuffix = line.substring(0, lengthNoSuffix);

      var trimmedLines = self.state.trimmedLines;
      if (line.length === lengthNoSuffix) {
        trimmedLines[index] = lineNoSuffix;
      } else {
        trimmedLines[index] = lineNoSuffix + suffix;
      }

      self.setState({
        trimmedLines: trimmedLines,
      });
    }
  }

  render () {
    var lines = [];
    if (this.props) {
      lines = this.props.text.split('\n').slice(0, this.props.lineCount);
      console.log(`lines:${lines}, length:${lines.length}`)
    }
    return (
      <div>
        {lines.map((line, index) => (
          <TrimToPx
            key={`trim-${index}`}
            text={line}
            suffix={this.props.suffix}
            width={this.props.width}
            trimmedLength={this.trimmedLength(line, index)}
          />
          )
        )}
        {this.state.trimmedLines.map((line, index) => (
          <span key={`span-${index}`} >{line}<br /></span>
          )
        )}
      </div>
      )
  }
}

export default withStyles(styles)(TextTruncate);
