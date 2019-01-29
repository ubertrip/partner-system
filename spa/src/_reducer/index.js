export const TOGGLE_LOADER = 'toggle-loader';
export const IS_AUTH = 'authorization';
export const IS_LOGOUT = 'logout';


const initialState = {
  title: 'Uber Trip',
  isLoading: false,
  loadingMessage: '',
  isAuth: false,
  islogout: true,
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

        case IS_LOGOUT:
        console.log('IS_LOGOUT', action);
        
        return{
          ...state,
          islogout: action.value,
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

export const islogout = (value = true) => ({
  type: IS_LOGOUT,
  value,
})