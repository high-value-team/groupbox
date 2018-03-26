import React  from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import TrimToPx from './TrimToPx';
import Typography from 'material-ui/Typography';


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

// export const calculateParts = (line) => {
//   return [
//     {text:"abc ", isURL:false, isLast:false},
//     {text:"https://www.google.com/", isURL:true, isLast:false},
//     {text:" hello", isURL:false, isLast:true},
//   ];
// };

export const calculateParts = (line) => {
  const parts = [];
  let current = 0;

  for (let i=0; i<line.URLPositions.length; i++) {
    const start = line.URLPositions[i][0];
    const end = line.URLPositions[i][1];

    if (current < start) {
      parts.push({
        text: line.originalText.substring(current, start),
        URL: '',
        isURL: false,
        isLast: false,
      });
    }

    if (end <= line.lengthNoSuffix) {
      parts.push({
        text: line.originalText.substring(start, end),
        URL: line.originalText.substring(start, end),
        isURL: true,
        isLast: false,
      });
      current = end;
    } else {
      parts.push({
        text: line.originalText.substring(start, line.lengthNoSuffix) + line.suffix,
        URL: line.originalText.substring(start, end),
        isURL: true,
        isLast: true,
      });
      current = line.lengthWithSuffix;
    }
  }

  // console.log(`line:
  // ${JSON.stringify(line, null, 2)}
  // originalText.length:${line.originalText.length}`);

  if (current <= line.lengthNoSuffix) {
    const suffix = line.lengthNoSuffix === line.originalText.length ? '' : line.suffix;
    parts.push({
      text: line.originalText.substring(current, line.lengthNoSuffix) + suffix,
      URL: '',
      isURL: false,
      isLast: true,
    });
  }

  console.log(`line:
  ${JSON.stringify(line, null, 2)}
  originalText.length:${line.originalText.length}`);


  return parts;
};

class TruncateText extends React.Component {
  static propTypes = {
    text: PropTypes.string.isRequired,
    suffix: PropTypes.string,
    width: PropTypes.number.isRequired,
    lineCount: PropTypes.number.isRequired,
  };

  constructor(props) {
    super(props);
    this.state = {
      lines: new Array(props.lineCount).fill({
        originalText: '',
        URLPositions: [],
        suffix: props.suffix,
        lengthWithSuffix: 0,
        lengthNoSuffix: 0,
        parts: [],
      }),
    };
    this.getSuffix = this.getSuffix.bind(this);
    this.trimmedLength = this.trimmedLength.bind(this);
    this.calculateLinks = this.calculateLinks.bind(this);
  }

  getSuffix() {
    return this.props.suffix ? this.props.suffix : '';
  }

  trimmedLength(originalText, index) {
    const suffix = this.getSuffix();
    const self = this;

    return (lengthWithSuffix) => {
      // console.log(`index:${index} length:${lengthWithSuffix}`);
      const lengthNoSuffix = lengthWithSuffix - suffix.length;

      const line = Object.assign({}, self.state.lines[index]);
      line.originalText = originalText;
      line.lengthWithSuffix = lengthWithSuffix;
      line.lengthNoSuffix = lengthNoSuffix;
      line.URLPositions = self.calculateLinks(originalText, lengthNoSuffix);
      line.parts = calculateParts(line);

      self.state.lines[index] = line;
      self.forceUpdate();
    };
  }

  calculateLinks(line, length) {
    // eslint-disable-next-line
    let regex = /((?:(http|https|Http|Https|rtsp|Rtsp):\/\/(?:(?:[a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(?:\%[a-fA-F0-9]{2})){1,64}(?:\:(?:[a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(?:\%[a-fA-F0-9]{2})){1,25})?\@)?)?((?:(?:[a-zA-Z0-9][a-zA-Z0-9\-]{0,64}\.)+(?:(?:aero|arpa|asia|a[cdefgilmnoqrstuwxz])|(?:biz|b[abdefghijmnorstvwyz])|(?:cat|com|coop|c[acdfghiklmnoruvxyz])|d[ejkmoz]|(?:edu|e[cegrstu])|f[ijkmor]|(?:gov|g[abdefghilmnpqrstuwy])|h[kmnrtu]|(?:info|int|i[delmnoqrst])|(?:jobs|j[emop])|k[eghimnrwyz]|l[abcikrstuvy]|(?:mil|mobi|museum|m[acdghklmnopqrstuvwxyz])|(?:name|net|n[acefgilopruz])|(?:org|om)|(?:pro|p[aefghklmnrstwy])|qa|r[eouw]|s[abcdeghijklmnortuvyz]|(?:tel|travel|t[cdfghjklmnoprtvwz])|u[agkmsyz]|v[aceginu]|w[fs]|y[etu]|z[amw]))|(?:(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9])\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[0-9])))(?:\:\d{1,5})?)(\/(?:(?:[a-zA-Z0-9\;\/\?\:\@\&\=\#\~\-\.\+\!\*\'\(\)\,\_])|(?:\%[a-fA-F0-9]{2}))*)?(?:\b|$)/gi;
    const URLPositions = [];
    let match;
    while ((match = regex.exec(line)) !== null) {
      const start = match.index;
      const end = regex.lastIndex;
      // console.log(`length:${length} start:${start} end:${end}`);
      if (start > length-1) {
        break;
      }
      URLPositions.push([start, end]);
    }
    return URLPositions;
  }

  render () {
    let textSplits = [];
    if (this.props) {
      textSplits = this.props.text.split('\n').slice(0, this.props.lineCount);
    }
    return (
      <div>
        {textSplits.map((text, index) => (
          <div key={`trim-${index}`} className={`trim-${index}`}>
            <TrimToPx
              key={`trim-${index}`}
              text={text}
              suffix={this.props.suffix}
              width={this.props.width}
              trimmedLength={this.trimmedLength(text, index)}
            />
          </div>
        ))}
        {this.state.lines.map((line, index) => (
          <Typography component="p" key={index}>
            <span>
              {line.parts.map((part, idx) =>
                part.isURL ? <a key={idx} href={part.text}>{part.text}</a> : part.text)}<br/>
            </span>
          </Typography>
        ))}
      </div>
    );
  }
}

export default withStyles(styles)(TruncateText);
