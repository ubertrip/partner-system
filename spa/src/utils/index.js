export const calcDriverSalary = data => {
  const {weeklyPayment, report} = data;
  return ((((weeklyPayment.netFares+weeklyPayment.miscPayment)*0.4)-report.diff) + (weeklyPayment.incentives*0.7)).toFixed(2);
};