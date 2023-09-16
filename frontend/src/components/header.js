import styles from '@/styles/header.module.css';
import Container from './container';

export default function Header() {
  return (
    <header className={styles.header}>
      <Container large>
        <div className={styles.flexContainer}>
          <div className={styles.header_title}>
            <h1>ウミガメのスープ</h1>
            <p>水平思考の間</p>
          </div>
          <ul className={styles.header_menu}>
            <li>ホーム</li>
            <li>問題一覧</li>
            <li>log in</li>
          </ul>
        </div>
      </Container>
    </header>
  );
}
