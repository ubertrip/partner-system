import React from 'react';
import moment from 'moment';

export const StatementSelect = props => <div className="statement-select">
  <select value={props.value} onChange={props.onChange}>
    {props.statements.map((s, i) => <option disabled={s.hidden} key={s.uuid}
                                            value={s.uuid}>{!i ? 'Текущая неделя: ' : ''} {moment(s.startAt).format('DD.MM.YYYY')} - {moment(s.endAt).format('DD.MM.YYYY')}</option>)}
  </select>
</div>;