
import './auth.scss';
import React  from 'react';
import {Link} from 'react-router-dom';  
// import axios from 'axios';

// const MenuForm = props => <div className="">
//     <div className="menu">
//         <div className="infinity-menu-node-container">
//         <label>auth</label>
        
//         <input  type="submit" size=""   value={props.menus} />

//         </div>


//     </div>
// </div>


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
                    {/* <MenuForm 
                menus={this.state.auth}
                /> */}
                </div> 
                
                })}
        </div>

    }
    
}