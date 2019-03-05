export const TOGGLE_LOADER = 'toggle-loader';

const initialState = {
  title: 'Uber Trip',
  isLoading: false,
  loadingMessage: '',
};

export default (state = initialState, action) => {
  switch (action.type) {
    case TOGGLE_LOADER:

      return {
        ...state,
        isLoading: action.value,
        loadingMessage: action.text,
      };

    default:
      return state
  }
}

export const toggleLoading = (value = false, text = '') => ({
  type: TOGGLE_LOADER,
  value,
  text,
});