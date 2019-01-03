export const TOGGLE_LOADER = 'toggle-loader';
export const IS_AUTH = 'authorization';


const initialState = {
  title: 'Uber Trip',
  isLoading: false,
  loadingMessage: '',
  isAuth: false,
};

export default (state = initialState, action) => {
  switch (action.type) {
    case TOGGLE_LOADER:

      return {
        ...state,
        isLoading: action.value,
        loadingMessage: action.text,
      };
      case IS_AUTH:
      console.log('IS_AUTH', action);
      

        return{
          ...state,
          isAuth: action.value,
          loadingMessage: action.text
        }

    default:
      return state
  }
}

export const toggleLoading = (value = false, text = '') => (
  {
  type: TOGGLE_LOADER,
  value,
  text,
});

export const isAuth = (value = false) => ({
  type: IS_AUTH,
  value,
})