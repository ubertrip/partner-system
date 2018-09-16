import React, {Component} from 'react';

export default class WeeklyDashboard extends Component {
  state = {
    netFares: 0,
    cashCollected: 0,
    netPayout: 0,
    totalPercentEarn: 0,
    incentives: 0,
    additional: 0,
    driversCount: 6,
    fixedWeeklyEarn: 0,
  };

  componentWillReceiveProps(nextProps) {
    this.calc(nextProps);
  }

  calc = (nextProps) => {
    const min = 1550;
    const days = 5;

    const fixedWeeklyEarn = min * days * this.state.driversCount;

    let netFares = 0;
    let cashCollected = 0;
    let netPayout = 0;
    let incentives = 0;
    let additional = 0;

    // b3928042-6466-4d32-ba89-d565a3cb0d57 partner uuid

    nextProps.payments.filter(p => p.driver.uuid !== 'b3928042-6466-4d32-ba89-d565a3cb0d57').forEach(p => {
      const miscPayment = p.weeklyPayment.miscPayment !== null ? parseFloat(p.weeklyPayment.miscPayment) : 0; // возврат денег
      additional += miscPayment;
      incentives += p.weeklyPayment.incentives !== null ? parseFloat(p.weeklyPayment.incentives) : 0; // бонусы

      netFares += p.weeklyPayment.netFares !== null ? parseFloat(p.weeklyPayment.netFares) : 0; // общая сумма
      netFares += miscPayment; // доп выплата плюсуем к основной

      cashCollected += p.weeklyPayment.cashCollected !== null ? parseFloat(p.weeklyPayment.cashCollected) : 0; // нал
      netPayout += p.weeklyPayment.netPayout !== null ? parseFloat(p.weeklyPayment.netPayout) : 0; // безнал
    });

    this.setState({
      netFares,
      cashCollected,
      netPayout,
      additional,
      incentives,
      totalPercentEarn: ((netFares + additional) / fixedWeeklyEarn * 100).toFixed(2),
      fixedWeeklyEarn,
    })

  };

  render() {
    return <div className="weekly-dashboard">
      <table>
        <tbody>
        <tr>
          <td>Водителей:</td>
          <td><b>{this.state.driversCount}</b></td>
        </tr>

        <tr>
          <td>Всего:</td>
          <td><b>{this.state.netFares.toFixed(2)}</b> грн</td>
        </tr>

        <tr>
          <td>Получено наличными:</td>
          <td><b>{this.state.cashCollected.toFixed(2)}</b> грн</td>
        </tr>

        <tr>
          <td>Чистая сумма оплаты:</td>
          <td><b>{this.state.netPayout.toFixed(2)}</b> грн</td>
        </tr>

        <tr>
          <td>Доп. выплаты, возвраты:</td>
          <td><b>{this.state.additional.toFixed(2)}</b> грн</td>
        </tr>

        <tr>
          <td>Промоакции:</td>
          <td><b>{this.state.incentives.toFixed(2)}</b> грн</td>
        </tr>

        <tr>
          <td>Прибыль партнера 60%:</td>
          <td><b style={{color: 'green'}}>{(this.state.netFares * 0.6 + this.state.incentives * 0.3).toFixed(2)}</b> грн
          </td>
        </tr>

        <tr>
          <td>Прибыль водителей 40%:</td>
          <td><b style={{color: 'red'}}>{(this.state.netFares * 0.4 + this.state.incentives * 0.7).toFixed(2)}</b> грн
          </td>
        </tr>
        </tbody>
      </table>

      <div>
        План выполнен на: <b
        className={`${this.state.totalPercentEarn < 70 ? 'progress-low' : ''} ${this.state.totalPercentEarn >= 70 && this.state.totalPercentEarn <= 89 ? 'progress-mid' : ''} ${ this.state.totalPercentEarn >= 90 ? 'progress-ok' : ''}`}>{this.state.totalPercentEarn}%</b>
        <progress value={this.state.totalPercentEarn} max="100">План {this.state.totalPercentEarn}%</progress>
        <br/>

        {this.state.fixedWeeklyEarn - this.state.netFares > 0 ?
          <i style={{color: 'red'}}>Недостача: {(this.state.fixedWeeklyEarn - this.state.netFares).toFixed(2)} грн<br/></i> : null}
        {this.state.fixedWeeklyEarn - this.state.netFares <= 0 ? <i style={{color: 'green'}}>Недостача: нет</i> : null}
      </div>
    </div>
  }
}