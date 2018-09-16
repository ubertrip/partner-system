import React, {Component} from 'react';
import {connect} from 'react-redux';
import {getDriver} from './reducer'


class CSearchByDriverID extends Component {
  state = {id: ''};

  setID = e => this.setState({id: e.target.value});

  search = e => {
    e.preventDefault();

    this.props.getDriver(this.state.id);

    return false;
  };

  render() {
    return <div className="serach-driver">
      <h3>Просмотр платежей</h3>
      <form onSubmit={this.search}>
        <input type="number" autoFocus onChange={this.setID} placeholder="Введите идентификатор"/><br/>
        <button type="submit">Найти</button>
      </form>

      <p>
        <i>Для поиска введите последние 5 цифр указанные на вашей топливной карте</i>
        <img width={250} src="/assets/card.jpg" alt="card"/>
      </p>
    </div>;

  }
}

const mapStateToProps = state => ({});

export default connect(
  mapStateToProps,
  {
    getDriver,
  }
)(CSearchByDriverID)