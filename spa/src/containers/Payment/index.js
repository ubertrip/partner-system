import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router-dom';
import {loadDriverPayments, addPayment} from './reducer'
import moment from "moment";
import {calcDriverSalary, weeklyEarnSum} from '../../utils';
import {StatementSelect} from '../../components/StatementSelect';
import {onChangeStatementUUID, loadStatements} from "../Payments/actions";

const DriverWeeklyPayments = props => <div className="driver-payments-list">
  {props.payments.length ? <div>
    <h3>Платежи</h3>
    <div className="driver-weekly-table">
      <table>
        <thead>
        <tr>
          <th>Дата</th>
          <th>Внесено</th>
          <th>UUID</th>
          <th>Баланс после платежа</th>
          <th>Наличные</th>
          <th>Оператор</th>
          <th>Газ/Бензин</th>
          <th>Комментарий</th>
        </tr>
        </thead>
        <tbody>
        {props.payments.map((p, i) => <tr key={`driver-payment-item-${i}`}>
          <td>{moment.utc(p.createdAt).format('DD.MM.YYYY HH:mm')}</td>
          <td><b style={{color: 'green'}}>₴{p.credit.toFixed(2)}</b></td>
          <td>{p.paymentUuid}</td>
          <td>{p.balance}</td>
          <td>{p.cashCollected}</td>
          <td>{p.createdBy}</td>
          <td>{p.gas}/{p.petrol}грн</td>
          <td>{p.extra ? p.extra : '-'}</td>
        </tr>)}
        </tbody>
      </table>
    </div>
  </div> : <i>Нет платежей за этот период</i>}
</div>;

const AddCredit = props => <div>
  <input type="number" autoFocus value={props.credit} onChange={props.onChange.bind(this, 'diff')} placeholder="Сумма"/><br/>
  <input type="number" value={props.gas} onChange={props.onChange.bind(this, 'gas')} placeholder="Газ"/><br/>
  <input type="number" value={props.petrol} onChange={props.onChange.bind(this, 'petrol')} placeholder="Бензин"/><br/>
  <textarea style={{marginTop: 10}} value={props.extra} onChange={props.onChange.bind(this, 'extra')} placeholder="Комментарий" cols="30" rows="5" />
  <br/>
  <button disabled={props.credit <= 0} onClick={props.addPayment}>Оплатить</button>
</div>;

class CPayment extends Component {
  state = {diff: 0, extra: '', gas: '', petrol: ''};

  constructor(props) {
    super(props);

    this.statementUUID = this.props.match.params.statementUUID;
    this.driverUUID = this.props.match.params.driverUUID;
    this.mode = this.props.match.params.mode;

    this.props.loadStatements().then(() => {
      this.props.onChangeStatementUUID(this.statementUUID);
    });

    this.props.loadDriverPayments(this.statementUUID, this.driverUUID).then(() => {
      this.setState({diff: this.props.report.diff > 0 ? this.props.report.diff.toFixed(2) : 0, extra: '', gas: '', petrol: ''});
    });
  }

  onChangePayment = (type, e) => {
    const value = e.target.value;
    this.setState({[type]: value})
  };

  addPayment = () => {
    if (this.state.diff < 1 || this.state.diff > 3000) {
      alert("Минимальная сумма 1 гривна, максимальная 3000 гривен");
      return;
    }

    if(this.state.gas > 1000) {
      alert("Максимальная стоимость газа 1000грн");
      return;
    }

    if(this.state.petrol > 500) {
      alert("Максимальная стоимость бензина 500грн");
      return;
    }

    if(!this.state.gas && !window.confirm("Добавить платеж без газа?")) {
      return;
    }

    window.confirm(`Внести платеж на сумму ${this.state.diff}грн?\nГаз: ${this.state.gas ? this.state.gas : 0}грн\nБензин: ${this.state.petrol ? this.state.petrol : 0}грн`) && this.props.addPayment(
      this.statementUUID,
      this.driverUUID,
      this.state.diff,
      this.state.extra,
      this.state.gas,
      this.state.petrol
    ).then(() => {
      this.setState({diff: 0, extra: '', gas: '', petrol: ''});
    })
  };

  isEditMode = () => this.mode === 'add';

  onChangeStatementUUID = e => {
    const value = e.target.value;

    this.props.loadDriverPayments(value, this.driverUUID).then(() => {
      this.props.onChangeStatementUUID(value, this.driverUUID, this.isEditMode());
      this.setState({diff: this.props.report.diff > 0 ? this.props.report.diff.toFixed(2) : 0, extra: ''});
    });
  };

  calcEarnPercent = () => (this.props.weeklyPayment.netFares/(weeklyEarnSum())*100);

  render() {
    const {driver, weeklyPayment, report} = this.props;
    return <div className="driver-payments">
      {this.props.driver ? <div>

        {this.isEditMode() ? <div><Link to="/payments">{'<<<'}платежи</Link></div> : null}

        <h3>Водитель: {this.props.driver.name}</h3>

        <StatementSelect
          statements={this.props.statements}
          onChange={this.onChangeStatementUUID}
          value={this.props.statementUUID}
        />

        <h4>Данные за
          период: {moment(this.props.statement.startAt).format('DD.MM.YYYY')} - {moment(this.props.statement.endAt).format('DD.MM.YYYY')}</h4>


        <h2>Ваш недельный план<br/>выполнен на ({this.calcEarnPercent() < 100 ? <span style={{color: 'red'}}>{this.calcEarnPercent().toFixed(2)}</span> : this.calcEarnPercent().toFixed(2)}%)</h2>

        <h2 style={{marginBottom: 0}}>Сумма к выплате: <b style={{color: 'violet'}}>₴{calcDriverSalary({
          weeklyPayment,
          report
        })}</b></h2>
        <small>* данная сумма примерный расчет выплаты и может быть изменена</small>

        <div style={{marginTop: 10}}>
          <img className="driver-payments__photo" src={this.props.driver.photo} alt={this.props.driver.name}/>
        </div>

        <div className="driver-weekly-table">
          <table>
            <tbody>
            <tr>
              <td>Обновлено:</td>
              <td>{moment(weeklyPayment.updatedAt).isValid() ? moment(weeklyPayment.updatedAt).format('DD.MM.YYYY HH:mm') : '-'}</td>
            </tr>
            <tr>
              <td>ID:</td>
              <td><b>{driver.id}</b></td>
            </tr>
            <tr>
              <td><i>Баланс:</i></td>
              <td>₴{report.balance.toFixed(2)}</td>
            </tr>
            <tr>
              <td><i>Разница:</i></td>
              <td>
                <div>
                  ₴{<b style={{color: report.diff >= 1 ? 'red' : 'green'}}>{report.diff.toFixed(2)}</b>}
                </div>
                <div>
                  {this.isEditMode() ? <AddCredit
                    onChange={this.onChangePayment}
                    addPayment={this.addPayment}
                    credit={this.state.diff}
                    extra={this.state.extra}
                    gas={this.state.gas}
                    petrol={this.state.petrol}
                  /> : null}
                </div>
              </td>
            </tr>
            <tr>
              <td>Тарифы без налогов и сборов:</td>
              <td>₴{weeklyPayment.netFares}</td>
            </tr>
            <tr>
              <td>Бонусы:</td>
              <td>₴{weeklyPayment.incentives}</td>
            </tr>
            <tr>
              <td>Платеж категории "Прочее":</td>
              <td>₴{weeklyPayment.miscPayment}</td>
            </tr>
            <tr>
              <td>Получено наличными:</td>
              <td>₴{weeklyPayment.cashCollected}</td>
            </tr>
            <tr>
              <td>Чистая сумма оплаты:</td>
              <td>₴{weeklyPayment.netPayout}</td>
            </tr>
            </tbody>
          </table>
        </div>

        <DriverWeeklyPayments
          payments={this.props.payments}
        />
      </div> : null}
    </div>;

  }
}

const mapStateToProps = state => ({
  driver: state.driverPayments.driver,
  payments: state.driverPayments.payments,
  report: state.driverPayments.report,
  weeklyPayment: state.driverPayments.weeklyPayment,
  statement: state.driverPayments.statement,

  statements: state.payments.statements,
  statementUUID: state.payments.statementUUID,
});

export default connect(
  mapStateToProps,
  {
    loadDriverPayments,
    addPayment,
    onChangeStatementUUID,
    loadStatements,
  }
)(CPayment)