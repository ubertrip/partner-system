
import './auth.scss';
import React  from 'react';


export default class Menu extends React.Component {
    render () {
        let menus = [
            "auth",
            "login",
            "driver",
            "payments"
        ]
        return <div>
            {menus.map((value, index)=>{
                //  return console.log(value)
                return <div key={index}><Link label={value} /></div>


            })}
        </div>
    }
}
class Link extends React.Component {
    render(){
    const url ="/" + this.props.label;
    return <div>        
        <a href={url}>{this.props.label}</a>
    </div>
    }
}
