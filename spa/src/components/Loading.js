import React from 'react';

export const Loading = props => <div className="loader">
  <div>Loading...</div>
  {props.text ? <div>{props.text}</div> : null}
</div>;