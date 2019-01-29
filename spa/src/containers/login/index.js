import React  from 'react';
// import PServer from './driver';
import { isAuth } from '../../_reducer'
import { connect } from 'react-redux';
import {auth} from '../Payment/reducer'
import { islogout } from '../../_reducer'
import axios from 'axios';
import {Redirect} from 'react-router';


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
  <form onSubmit={props.onlogout}>
  <div>
  <input type="submit" size="70" value="Выйти" />
  </div>
  </form>
  
  </div>


 class Login extends React.Component {
  constructor(props){
    super(props)
    this.state = {login: '123', password: '123', isLoading: false, loadingMessage: '', router: null};
  }

  setlogin = e => this.setState({login: e.target.value});

	onSendForm = e => {
    e.preventDefault();
    console.log("onSendForm")

    if(!this.state.login || !this.state.password) {
      alert('Заполните все поля формы');      
    	return;
    }

    // в этом месте отправляем на сервер    
    this.props.auth(this.state);
    //this.props.getUser(this.state.login);

    return false;
  };

  // console.log(error);
  // return alert("Cannot found user");

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
        logout={this.onlogout}
        // logout={this.logout}
      />
      {this.state.router}
    </div>;
  }

};
  class Logout extends React.Component {
    constructor(props){
      super(props)
      this.onClick=this.onlogout.bind(this);
    }
  
    // logoutAccaunt = e => {
    //   this.setState({login: e.target.value});
    // }
  
    onlogout = e => {
      e.preventDefault();
      this.props.logout();
      console.log("logout")
  
      const user = {
        login : this.state.login,
      }
      
  
    axios.get(`http://localhost:4321/login`,{user})
    .then(res => {
        console.log(res);
        console.log(res.data);
        // axios.defaults.headers.common['Authorization'] = null
        // doNextThing()
    })
// };
    // .then(({data}) => {
    //   if (data.status === 'ok') {
    //     console.log('if is works', data);
    //     // dispatch(islogout(true));
    //     this.setState({
    //       router: <Redirect to= "/logout" push />
    //     });
    // this.props.out(this.state);
    //   }
    // })
    
    .catch((error) => {
      alert('you are trespassing')
    });
  
    };
    logout = e => {
      localStorage.removeItem('authorization');
      console.log('logout');
      return false;
    };
  
  render() {
    const { user } = this.props;
        if (user === null) return <Redirect to='/login'/>;
    return <div> 
    <button onClick={this.onlogout}> Выход</button>
    </div>;
    
    // <form onClick={this.onlogout}
    // <input type="submit" />
    // </form>
    // <div style={style}>
    // </div>
    
    // 
    // <div className="wrapper serach-driver">
    /* <h1 className="wrapper">Выйти в UberTrip</h1>
    <LoginForm 
      logout={this.onlogout}

    /> */
   
}
  
  };

const mapStateToProps = state =>({
  logout: state.global.logout,
  state
});

const mapDispatchToProps = {
  islogout,
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
