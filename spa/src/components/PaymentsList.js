import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import './PaymentList.scss';
import moment from 'moment';
import {calcDriverSalary} from '../utils';

export default class PaymentsList extends Component {
  checkUptime = d => {
    const f = 'DD/MM/YYYY HH:mm';
    const start = moment(moment.utc(d).format(f), f);
    const end = moment(moment().format(f), f);

    const duration = moment.duration(end.diff(start));
    return duration.asMinutes() > 1 ? <span style={{color: 'red'}}>Expired: {start.format('DD.MM.YYYY HH:mm')}</span> : <span>{start.format('DD.MM.YYYY HH:mm')}</span>;
  };

  render() {
    return <div className="payments-list">
      <table>
        <thead>
        <tr>
          <th colSpan={2}>Водитель</th>
          <th>Обновлено</th>
          <th>ID</th>
          <th><i>Баланс</i></th>
          <th colSpan={2}><i>Разница</i></th>
          <th>Зарплата</th>
          <th>Тарифы без налогов и сборов</th>
          <th>Бонусы</th>
          <th>Платеж категории "Прочее"</th>
          <th>Получено наличными</th>
          <th>Чистая сумма оплаты</th>
        </tr>
        </thead>
        <tbody>
        {this.props.payments.map((p, i) => <tr className="payment-item" key={`payment-tr-${i}`}>
          <td><img src={p.driver.photo} alt=""/></td>
          <td><Link to={`/credit/${this.props.statementUUID}/${p.driver.uuid}/add`}>{p.driver.name}</Link></td>
          <td>{this.checkUptime(p.weeklyPayment.updatedAt)}</td>
          <td><b>{p.driver.id}</b></td>
          <td>₴{p.report.balance}</td>
          <td>₴{p.report.diff >= 1 ? <b style={{color: 'red'}}>{p.report.diff.toFixed(2)}</b> : <b style={{color: 'green'}}>{p.report.diff.toFixed(2)}</b> }</td>
          <td>{p.report.diff >= 1 ? <Link to={`/credit/${this.props.statementUUID}/${p.driver.uuid}/add`}>Оплатить</Link> : null}</td>
          <td style={{color: 'violet'}}>₴{Math.round(calcDriverSalary(p))}</td>
          <td>₴{p.weeklyPayment.netFares}</td>
          <td>₴{p.weeklyPayment.incentives}</td>
          <td>₴{p.weeklyPayment.miscPayment}</td>
          <td>₴{p.weeklyPayment.cashCollected}</td>
          <td>₴{p.weeklyPayment.netPayout}</td>
        </tr>)}
        </tbody>
      </table>
    </div>;

  }
}