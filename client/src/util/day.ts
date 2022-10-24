import dayjs from 'dayjs'

export const parseDay = (day: string) => {
  return dayjs(new Date(day)).format('YYYY/MM/DD')
}
