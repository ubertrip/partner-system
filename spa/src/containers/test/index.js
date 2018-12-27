import React  from 'react';
import './index.css';


function  PlayButton(props) {
    const className = props.isMusicPlaying ? 'play active' : 'play';
    
    return (     
        <h1>Hello, I am a React Component!</h1>,
        <a onClick={props.onClick} href="#" title="Play video" className={className} />
    );
  };

function  Pictures(props) {
const className = props.isMusicPlaying ? 'play active' :'hidden';
return (     
    <h1>Hello, I am a React Component!</h1>,
    <a onClick={props.onClick} href="#" title="Play" className={className} />
);
};

  export default class test extends React.Component {
    constructor(props) {
        super(props);
        this.state = { isMusicPlaying: false };
      }
      handleClick() {
        if (this.state.isMusicPlaying) {
            this.setState({ isMusicPlaying: false });
        } else {
            this.setState({ isMusicPlaying: true });
        }
      };

    render() {
        // let status = this.state.isMusicPlaying ? 'Playing' : 'Not playing';
        return (
            <div>
                {/* <h1 onClick={this.handleClick.bind(this)}>{ status } </h1> */}
               
                < PlayButton
                onClick={this.handleClick.bind(this)}
                isMusicPlaying={this.state.isMusicPlaying}  />
                
                < Pictures
                onClick={this.handleClick.bind(this)}
                isMusicPlaying={this.state.isMusicPlaying}  />
               
            </div>
        );

    }
}

