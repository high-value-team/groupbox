import * as React from 'react';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Typography from 'material-ui/Typography';
import IconButton from 'material-ui/IconButton';
import MenuIcon from 'material-ui-icons/Menu';
import withStyles from 'material-ui/styles/withStyles';

const decorate = withStyles(() => ({
  root: {
    width: '100%',
  },
  flex: {
    flex: 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
}));

interface Props {
  version: string;
}

export default decorate<Props>(({ classes, version }) => (
  <div className={classes.root}>
    <AppBar position="static">
      <Toolbar>
        <IconButton className={classes.menuButton} color="contrast" aria-label="Menu">
          <MenuIcon />
        </IconButton>
        <Typography type="title" color="inherit" className={classes.flex}>
          Groupbox {version}
        </Typography>
      </Toolbar>
    </AppBar>
  </div>
));
