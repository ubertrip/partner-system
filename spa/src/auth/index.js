import React  from 'react';
import {Link} from 'react-router-dom';

export default class Menu extends React.Component {
    constructor(props){
    super(props)
    this.state = {
         menus : [
            "auth",
            "login",
            "driver",
            "payments"
        ]}
    }
        
        render () {
        return <div className="menu">
            {this.state.menus.map((value, index)=>{
                return <div key={index}>
                    <Link to={value}>{value}</Link>
                </div> 
                })}
        </div>
    }
}