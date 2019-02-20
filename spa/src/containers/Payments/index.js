import React, {Component} from 'react';
import {connect} from 'react-redux';
import {loadPayments, loadStatements, onChangeStatementUUID, changeStatement} from './actions'
import PaymentsList from '../../components/PaymentsList'
import WeeklyDashboard from '../../components/WeeklyDashboard'
import {StatementSelect} from '../../components/StatementSelect';

class Index extends Component {
  constructor(props) {
    super(props);
    this.props.loadStatements().then(() => {
      this.props.onChangeStatementUUID(this.props.statements && this.props.statements.length ? this.props.statements[0].uuid : '');
      this.props.loadPayments(this.props.statementUUID);
    });
  }

  onChangeStatementUUID = e => {
    const value = e.target.value;
    this.props.changeStatement(value);
  };

  render() {
    // console.log('in paymets,this.props.isAuth',this.props.isAuth);
    return <div>
      <div style={{padding: '10px 5px'}}>
        <StatementSelect
          statements={this.props.statements}
          onChange={this.onChangeStatementUUID}
          value={this.props.statementUUID}
        />
      </div>

      <WeeklyDashboard
        payments={this.props.payments}
      />

      <PaymentsList
        payments={this.props.payments}
        statementUUID={this.props.statementUUID}
      />
    </div>;

  }
}

const mapStateToProps = state => ({
  payments: state.payments.list,
  statements: state.payments.statements,
  statementUUID: state.payments.statementUUID,
  // isAuth: state.global.isAuth,

});

export default connect(
  mapStateToProps,
  {
    loadPayments,
    loadStatements,
    changeStatement,
    onChangeStatementUUID,
  }
)(Index)