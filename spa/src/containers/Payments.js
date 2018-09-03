import React, {Component} from 'react';
import {connect} from 'react-redux';

class Payments extends Component {
  render() {
    return <div>payments</div>;

  }
}

const mapStateToProps = state => ({});

export default connect(
  mapStateToProps,
  null
)(Payments)