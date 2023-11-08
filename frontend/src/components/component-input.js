import styles from '@/styles/component-input.module.css';

export default function MyInput({ label, type, size, value, onChange }) {
  return (
    <div className={styles.input_wrapper}>
      <label className={styles.input_label}>{label}</label>
      <input className={styles.input} size={size} type={type} value={value} onChange={onChange} />
    </div>
  );
}
