import React  from 'react';
import { connect } from 'react-redux';
import {auth} from '../Payment/reducer'
import {toggleLoading} from '../../_reducer';
import {logout} from '../Payment/reducer';


const LoginForm = props => <div className="wrapper serach-driver">
  
  <form onSubmit={props.sendForm}>
  <fieldset className="wrapper-login">  
    <div className="wrapper-login">
      <label>Логин:</label>
      <input  type="text" size="70px" value={props.login} onChange={props.setProp.bind(this, 'login')} />
    </div>
    
    <div className="wrapper-login">
      <label>Пароль:</label>
      <input type="password" size="70px" value={props.password} onChange={props.setProp.bind(this, 'password')}/>
    </div>
    </fieldset>
    <div>
      <input type="submit" value="Войти" />
    </div>
  </form>
  </div>

 class Login extends React.Component {
  constructor(props){
    super(props)
    this.state = {login: '123', password: '123', isLoading: false, loadingMessage: '', router: null};
  }

  // setlogin = e => this.setState({login: e.target.value});

	onSendForm = e => {
    e.preventDefault();
    console.log("onSendForm" )

    if(!this.state.login || !this.state.password) {
      alert('Заполните все поля формы');      
    	return;
    }
    // в этом месте отправляем на сервер    
    this.props.auth(this.state);
    //this.props.getUser(this.state.login);
    toggleLoading(false);
    return false;
  };

  onSetProp = (prop, e) => {
  	const value = e.target.value;
    this.setState({[prop]: value});
  };
  

  render() {
    return <div className="wrapper serach-driver">
      <h1 className="wrapper">Войти в UberTrip</h1>
      <LoginForm 
        login={this.state.login} 
        password={this.state.password} 
        sendForm={this.onSendForm}
        setProp={this.onSetProp}
        logout={this.props.onLogout}
      />
      <Logout logout={logout} ref="logout" />
      {this.state.router}
    </div>;
  }

};

class Logout extends React.Component {

  onLogout() {
    console.log("logout")
    this.props.logout(this.state);
  };

  render() {
    return <div className="button"> 
      <button onClick={e => this.onLogout(e)}>Выход</button>
    </div>
    }
};

export default connect(
  null, 
  {
    auth,
  }
  )(Login);

connect(
  null,
  {
    logout
  }
  )(Logout);
  
  connect(
    null,
    {
      logout
    }
    )(Logout);
