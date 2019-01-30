import React  from 'react';
// import PServer from './driver';
import { isAuth } from '../../_reducer'
import { connect } from 'react-redux';
import {auth} from '../Payment/reducer'
import {logout} from '../Payment/reducer'
// import { islogout } from '../../_reducer'


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

    <div>
      <input type="checkbox" name="remember"/>Запомнить меня
    </div>
    </fieldset>
    <div>
      <input type="submit" size="70px" value="Войти" />
    </div>
  </form>
  {/* <form onSubmit={props.logout}>
  <div>
  <input type="submit" size="70" value="Выйти" />
  </div>
  </form> */}
  {/* <button type="button" onClick={props.logout}>Выйти</button> */}
  
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
        logout={this.props.onlogout}
      />
      <button type="button"  onClick={this.logout}>Выход</button>
      {this.state.router}
      {/* {this.state.onlogout} */}
    </div>;
  }

};
class Logout extends React.Component {


  onlogout = e => {
    e.preventDefault();
    console.log("logout")

     this.onClick=this.onlogout.bind(this);

    this.props.logout(this.state);
    
    return false;
  };

  render() {
    return 
    }
};
  // constructor(props){
  //   super(props)
  //   this.state = {status: "ok",}
    // this.onClick=this.onlogout.bind(this);
  // }

  // logoutAccaunt = e => {
  //   this.setState({login: e.target.value});
  // }

  
    // e.preventDefault();
    // this.props.logout();
    
  //   return PaymentsApi.logout(login, password).then(({data}) => {
  //   if(!data.status === 'ok') {
  //     console.log("logout", data)
  //     // return;
  //   }
  // })
  //   // const user = {
    //   login : this.state.login,
    // }
    
  //   return false;
  // };
  
  // logout = e => {
  //   localStorage.removeItem('authorization');
  //   console.log('logout');
  //   return false;
  // };


const mapStateToProps = state =>({
  logout: state.global.logout,
  state
});

const mapDispatchToProps = {
  logout,
};

export default connect(
  null, 
  {
    isAuth,
    auth,
    // islogout,
  }
  )(Login)

connect(mapStateToProps, mapDispatchToProps)(Logout);
