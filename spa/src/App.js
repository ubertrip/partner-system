import React, {Component} from 'react';
import {Router, Switch, Route} from 'react-router-dom';
import {connect} from 'react-redux';
import {history} from './store';
import './App.scss';
import Payments from './containers/Payments';
import BaseLayout from './containers/BaseLayout';

const Index = props => <div style={{width: '800px', margin: '0 auto', textAlign: 'center'}}>
  <img src="/assets/trip.jpg" alt=""/>
  <h2>+38-050-551-62-60</h2>
</div>;

class App extends Component {
  render() {
    return <Router history={history}>
      <BaseLayout>
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
