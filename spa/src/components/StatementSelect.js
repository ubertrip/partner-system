import React from 'react';

export const StatementSelect = props => <div>
  <select value={props.value} onChange={props.onChange}>
    {props.statements.map((s, i) => <option disabled={s.hidden} key={s.uuid}
                                            value={s.uuid}>{!i ? 'Current' : ''} {s.startAt}  - {s.endAt}</option>)}
  </select>
</div>;