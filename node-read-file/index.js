const fs = require('fs');
const path = require('path');

/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
exports.listFiles = (req, res) => {
  // カレントディレクトリを読み込み
  (function () {
    currentDir = process.cwd();
    console.log(
      JSON.stringify({
        severity: 'DEBUG',
        message: { currentDir },
      }),
    );
  })();

  // root 配下のディレクトリを読み込み
  (function () {
    rootDirs = fs.readdirSync('/');
    console.log(
      JSON.stringify({
        severity: 'DEBUG',
        message: { rootDirs },
      }),
    );
  })();

  // カレントディレクトリ配下のファイルを再帰的に読み込み
  (function () {
    console.log(
      JSON.stringify({
        severity: 'DEBUG',
        message: {
          files: fs
            .readdirSync('./', { withFileTypes: true })
            .map((dir) => dir.name),
        },
      }),
    );
  })();

  res.status(200).send('done');
};
