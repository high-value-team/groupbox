import TextTruncate from '../TextTruncate';

export default {
  component: TextTruncate,
  props: {
    text: "abc123 - Groupbox \nhallo\n welt \nehllo awesome world out there",
    width: 500,
    suffix: "...",
    lineCount: 3,
  }
};
