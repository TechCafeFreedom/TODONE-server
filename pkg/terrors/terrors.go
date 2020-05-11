package terrors

import (
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

// TodoneError サーバ-クライアント間エラーハンドリング用エラー
type TodoneError struct {
	// エラーコード
	ErrorCode int
	// システムエラーメッセージ(日本語)
	ErrorMessageJP string
	// システムエラーメッセージ(英語)
	ErrorMessageEN string
	// xerrors拡張用フィールド
	err error
	// それぞれでfmt.Errorf("%w", err)を記述する必要があるためxerrors使う。
	frame xerrors.Frame
}

// New TodoneErrorを生成する
func New(errorCode int) error {
	return newError(nil, errorCode, "", "")
}

// Newf TodoneErrorをエラーメッセージ付きで生成する
func Newf(errorCode int, messageJP string, messageEN string) error {
	return newError(nil, errorCode, messageJP, messageEN)
}

// Wrap エラーをTodoneエラーでラップする
func Wrap(cause error, errorCode int) error {
	return newError(cause, errorCode, "", "")
}

// Wrapf エラーをTodoneエラーで、エラーメッセージ付きでラップする
func Wrapf(cause error, errorCode int, messageJP, messageEN string) error {
	return newError(cause, errorCode, messageJP, messageEN)
}

func newError(cause error, errorCode int, errorMessageJP, errorMessageEN string) error {
	return &TodoneError{
		ErrorCode:      errorCode,
		ErrorMessageJP: errorMessageJP,
		ErrorMessageEN: errorMessageEN,
		err:            cause,
		frame:          xerrors.Caller(2),
	}
}

// Stack エラーをStackする
// スタックフレームを明示的に積んでいく必要があるためエラー出力に記録したいエラーハンドリング箇所ではStackを行う
func Stack(err error) error {
	var errorCode int
	var errorMessageJP, errorMessageEN string
	var todoneError *TodoneError
	if ok := xerrors.As(err, &todoneError); ok {
		errorCode = todoneError.ErrorCode
		errorMessageJP = todoneError.ErrorMessageJP
		errorMessageEN = todoneError.ErrorMessageEN
	} else {
		return &TodoneError{
			ErrorCode:      http.StatusInternalServerError,
			ErrorMessageJP: "エラーのコンバート時にエラーが発生しました",
			ErrorMessageEN: "Error occured at covert to original error",
			err:            err,
			frame:          xerrors.Caller(1),
		}
	}
	return &TodoneError{
		ErrorCode:      errorCode,
		ErrorMessageJP: errorMessageJP,
		ErrorMessageEN: errorMessageEN,
		err:            err,
		frame:          xerrors.Caller(1),
	}
}

// Error エラーメッセージを取得する
func (e *TodoneError) Error() string {
	return fmt.Sprintf("messageJP=%s, messageEN=%s", e.ErrorMessageJP, e.ErrorMessageEN)
}

func (e *TodoneError) Unwrap() error {
	return e.err
}

func (e *TodoneError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *TodoneError) FormatError(p xerrors.Printer) error {
	p.Print(e.ErrorMessageJP, e.ErrorMessageEN)
	e.frame.Format(p)
	return e.err
}
