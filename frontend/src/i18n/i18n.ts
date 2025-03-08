import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import en from './en/common.json';

i18n
    .use(initReactI18next)
    .init({
        resources: {
            en: { translation: en },
        },
        lng: 'en', // default language
        fallbackLng: 'en', // fallback if language not found
        interpolation: { escapeValue: false } // React already escapes
    });

export default i18n;