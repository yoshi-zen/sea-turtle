import styles from '@/app/questions/page.module.css';
import Layout12 from '@/components/layout-12columns';

export default function Questions() {
  return (
    <Layout12>
      <div className={styles.left_container}>
        <h1 className={styles.pageTitle}>問題一覧</h1>
        <div className={styles.q1}>
          <p>1. ウミガメのスープ</p>
        </div>
        <div className={styles.q2}>
          <p>2. クレイジータクシー</p>
        </div>
        <div className={styles.q3}>
          <p>3. ひどいスッピン</p>
        </div>
        <div className={styles.q4}>
          <p>4. ゴンとチロ</p>
        </div>
        <div className={styles.q4}>
          <p>5. 過激なダイエット</p>
        </div>
        <div className={styles.q4}>
          <p>6. 爆弾が仕掛けられた車</p>
        </div>
        <div className={styles.q4}>
          <p>7. ある偉人の版画</p>
        </div>
        <div className={styles.q4}>
          <p>8. 値段より大切なもの</p>
        </div>
        <div className={styles.q4}>
          <p>9. 包帯を巻いたプログラマー</p>
        </div>
        <div className={styles.q4}>
          <p>10. 刈り取り</p>
        </div>
        <div className={styles.q4}>
          <p>11. 知る人ぞ知るラーメン</p>
        </div>
        <div className={styles.q4}>
          <p>12. 男の覚悟</p>
        </div>
        <div className={styles.q4}>
          <p>13. 道を教えてください</p>
        </div>
        <div className={styles.q4}>
          <p>14. 振られましたね？</p>
        </div>
        <div className={styles.q4}>
          <p>15. 真逆の主人公</p>
        </div>
        <div className={styles.q4}>
          <p>16. プリクラ開発秘話</p>
        </div>
        <div className={styles.q4}>
          <p>17. カメラマンの願い</p>
        </div>
        <div className={styles.q4}>
          <p>18. 臓器提供</p>
        </div>
        <div className={styles.q4}>
          <p>19. 素敵なステーキ</p>
        </div>
        <div className={styles.q4}>
          <p>20. うんこを見ろ</p>
        </div>
        <div className={styles.q4}>
          <p>21. 親を食った男</p>
        </div>
        <div className={styles.q4}>
          <p>22. 箱の中</p>
        </div>
        <div className={styles.q4}>
          <p>23. テレビ壊しちゃった</p>
        </div>
        <div className={styles.q4}>
          <p>24. パーカーを着なさい</p>
        </div>
      </div>
      <div className={styles.center_container}>
        <div className={styles.flex_wrapper}>
          <div className={styles.question_inner}>
            <p>「ウミガメのスープが飲みたい。」</p>
            <p>男は各地を旅行した。</p>
            <p>とある海辺のレストランで ウミガメのスープを出す店を見つけた。</p>
            <p>彼はウミガメのスープを注文し その日に自殺した。</p>
            <p>念願だったウミガメのスープが見つかったのに 彼はなぜ自殺したのだろうか？</p>
          </div>
          <div>
            <p>ここにチャット</p>
          </div>
        </div>
      </div>
    </Layout12>
  );
}
