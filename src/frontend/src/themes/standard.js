import { createMuiTheme } from 'material-ui/styles';

/**
 * To see what are the values you can override, use the source:
 * https://material-ui-next.com/customization/themes/
 */

const standardTheme = createMuiTheme({
  fontFamily: 'Roboto, sans-serif',
  palette: {
    type: 'light',
  },
});

export default standardTheme;
