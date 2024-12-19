export const formatDateRange = (start, stop) => {
  const startDate = new Date(start);
  const stopDate = new Date(stop);

  const optionsTime = { hour: "2-digit", minute: "2-digit" };
  const optionsFull = { month: "short", day: "numeric", hour: "2-digit", minute: "2-digit" };

  const isSameDay =
    startDate.getFullYear() === stopDate.getFullYear() &&
    startDate.getMonth() === stopDate.getMonth() &&
    startDate.getDate() === stopDate.getDate();

  const formatter = new Intl.DateTimeFormat("en-US", isSameDay ? optionsTime : optionsFull);

  const startFormatted = new Intl.DateTimeFormat("en-US", optionsFull).format(startDate);
  const stopFormatted = formatter.format(stopDate);

  return isSameDay ? `${startFormatted} - ${stopFormatted}` : `${startFormatted} - ${stopFormatted}`;
};
