import styles from '@/app/questions/page.module.css';
import Layout12 from '@/components/layout-12columns';

export default function Questions() {
  return (
    <Layout12>
      <div className={styles.left_container}>
        <h1 className={styles.pageTitle}>問題一覧</h1>
        <div className={styles.q1}>
          <p>1. ウミガメ</p>
        </div>
        <div className={styles.q2}>
          <p>aaa</p>
        </div>
        <div className={styles.q3}>
          <p>aaa</p>
        </div>
        <div className={styles.q4}>
          <p>aaa</p>
        </div>
      </div>
    </Layout12>
  );
}
