import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';

import BoxService from '../services/box';

import { withStyles } from 'material-ui/styles';
import { LinearProgress } from 'material-ui/Progress';

import Card, { CardHeader, CardContent, CardActions } from 'material-ui/Card';
import Avatar from 'material-ui/Avatar';
import IconButton from 'material-ui/IconButton';
import Typography from 'material-ui/Typography';
import FavoriteIcon from 'material-ui-icons/Favorite';
import Tooltip from 'material-ui/Tooltip';
import red from 'material-ui/colors/red';


const styles = theme => ({
  root: {
    width: '100%',
  },
  card: {
    float: 'left',
    margin: '0.5em',
    maxWidth: 400,
  },
  media: {
    height: 194,
  },
  expand: {
    transform: 'rotate(0deg)',
    transition: theme.transitions.create('transform', {
      duration: theme.transitions.duration.shortest,
    }),
  },
  expandOpen: {
    transform: 'rotate(180deg)',
  },
  avatar: {
    backgroundColor: red[500],
  },
  flexGrow: {
    flex: '1 1 auto',
  },
});

class Box extends React.Component {

  static propTypes = {
    classes: PropTypes.object.isRequired,
    boxkey: PropTypes.string,
  }

  constructor(props) {
    super(props);
    this.state = {
      status: 0, // 0: none, 1: loading, 2: loaded - ok, 3: loaded - error
      box: null,
    };
  }

  componentDidMount() {
    this.loadBox();
  }

  componentWillUnmount() {
    BoxService.next(null);
  }


  loadBox() {
    this.setState({ status: 1 });
    fetch(`/api/boxes/${this.props.boxkey}`).then(rsp => {
      if (rsp.ok) {
        rsp.json().then(box => {
          BoxService.next(box);
          this.setState({ box, status: 2 });
        });
      } else {
        this.setState({ status: 3 });
      }
    });

  }

  renderBlank() {
    return <div className={this.props.classes.root} />;
  }

  renderLoading() {
    return (
      <div className={this.props.classes.root}>
        <LinearProgress />
      </div>
    );
  }

  renderLoadedError() {
    return (
      <div className={this.props.classes.root}>
        <p>Die Box kann leider nicht gefunden werden.</p>
      </div>
    );
  }

  renderItem(index, item) {
    const { classes } = this.props;

    const authorInitials = item.authorNickname
      .split(/\s+/)
      .map(word => word.substr(0, 1))
      .reduce((initials, letter) => `${initials}${letter}`, '');

    return (
      <Card className={classes.card} key={`item-${index}`} >
        <CardHeader
          avatar={
            <Tooltip id="avatar-tooltip" title={item.authorNickname} placement="right">
              <Avatar aria-label="Author" className={classes.avatar}>
                {authorInitials}
              </Avatar>
            </Tooltip>
          }
          title={item.subject}
          subheader={moment(item.creationDate).fromNow()}
        />
        <CardContent>
          <Typography component="p">
            {item.message}
          </Typography>
        </CardContent>
        <CardActions disableActionSpacing={true}>
          <Tooltip id="favorite-tooltip" title="Finde ich gut!" placement="right">
            <IconButton aria-label="Finde ich gut!">
              <FavoriteIcon />
            </IconButton>
          </Tooltip>
          <div className={classes.flexGrow} />
        </CardActions>
      </Card>
    );
  }

  renderLoadedOK() {

    const { classes } = this.props;
    const { box } = this.state;

    const items = box.items.sort((a, b) => b.creationDate > a.creationDate );

    return (
      <div className={classes.root}>
        {items.map((item, index) => this.renderItem(index, item))}
      </div>
    );

  }

  render() {

    switch (this.state.status) {
      case 1:
        return this.renderLoading();
      case 2:
        return this.renderLoadedOK();
      case 3:
        return this.renderLoadedError();
    }

    return this.renderBlank();

  }

}

export default withStyles(styles)(Box);
