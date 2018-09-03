import React, {Component} from 'react';
import {Router, Switch, Route, Link} from 'react-router-dom';
import {connect} from 'react-redux';
import {history} from './store';
import './App.scss';
import Payments from './containers/Payments';
import BaseLayout from './containers/BaseLayout';

const Index = props => <h2>{props.title}</h2>;

class App extends Component {
  render() {
    return <Router history={history}>
      <BaseLayout>
        <Link to="/">Home</Link>
        <Switch>
          <Route exact path="/" render={() => <Index {...this.props}/>}/>
          <Route exact path="/payments" component={Payments}/>
        </Switch>
      </BaseLayout>
    </Router>;
  }
}

const mapStateToProps = state => ({
  title: state.global.title,
});

export default connect(
  mapStateToProps
)(App)
