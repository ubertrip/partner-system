import {withRouter} from "react-router-dom";
import {Component} from 'react';
import {connect} from 'react-redux';

class BaseLayout extends Component {
  render() {
    return this.props.children;
  }
}

const mapStateToProps = state => ({});

export default withRouter(connect(
  mapStateToProps,
  null
)(BaseLayout))