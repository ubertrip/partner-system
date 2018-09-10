const env = process.env.UBERTRIP_ENV || 'dev';

console.log('env', env);

export default require(`./${env}`).default;