import TruncateText from '../TruncateText';

export default {
  component: TruncateText,
  props: {
    // text: "abc123 - Groupbox \nhallo\n welt \nehllo awesome world out there",
    text: "hello world http://www.google.com this isabc https://florian.fnbk.cc/abc huhu",
    width: 361,
    suffix: "...",
    lineCount: 3,
  }
};
