import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Loading} from '../components/Loading';

class CLoading extends Component {
  render() {
    return this.props.isLoading ?  <Loading text={this.props.loadingMessage} /> : null;
  }
}

const mapStateToProps = state => ({
  isLoading: state.global.isLoading,
  loadingMessage: state.global.loadingMessage,
});

export default connect(
  mapStateToProps, null
)(CLoading)