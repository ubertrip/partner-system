import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import './PaymentList.scss';

export default class PaymentsList extends Component {
  render() {
    return <div>
      <table style={{width: '100%'}}>
        <tbody>
        {this.props.payments.map((p, i) => <tr className="payment-item" key={`payment-tr-${i}`}>
          <td><img src={p.driver.photo} alt=""/></td>
          <td><Link to={`/credit/${this.props.statementUUID}/${p.driver.uuid}`}>{p.driver.name}</Link></td>
          <td>{p.weeklyPayment.netFares}</td>
          <td>{p.weeklyPayment.incentives}</td>
          <td>{p.weeklyPayment.miscPayment}</td>
          <td>{p.weeklyPayment.netPayout}</td>
          <td>{p.weeklyPayment.cashCollected}</td>
        </tr>)}
        </tbody>
      </table>
    </div>;

  }
}