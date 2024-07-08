export const dotNotation = (key, obj) => {
  return key.split(".").reduce((o, i) => o[i], obj);
};

export const debounce = (func, wait) => {
  let timeout;

  return function (...args) {
    const context = this;

    clearTimeout(timeout);
    timeout = setTimeout(() => {
      func.apply(context, args);
    }, wait);
  };
};

export const truncate = (input, limit, separator = " ", indicator = "...") => {
  let result = "";
  let remainingLimit = limit;

  const strings = Array.isArray(input) ? input : [input];

  for (let str of strings) {
    if (!str) continue;
    if (result.length + str.length <= remainingLimit) {
      if (result) {
        result += separator;
      }
      result += str;
      remainingLimit -= str.length + (result ? separator.length : 0);
    } else {
      const cutIndex = remainingLimit - indicator.length;
      if (cutIndex > 0) {
        result +=
          (result ? separator : "") + str.substring(0, cutIndex) + indicator;
      } else {
        result += indicator;
      }
      break;
    }
  }

  return result;
};
