// 编辑器的状态
export enum Mode {
  ECH0 = 0, // 默认编辑状态
  Panel = 1, // 显示面板状态
  TODO = 2, // 待办事项状态
  EXTEN = 3, // 处理扩展状态
  PlayMusic = 4, // 音乐播放器状态
  Image = 5, // 图片上传状态
  TagManage = 6, // 标签管理状态
}

// 扩展类型
export enum ExtensionType {
  MUSIC = 'MUSIC',
  VIDEO = 'VIDEO',
  GITHUBPROJ = 'GITHUBPROJ',
  WEBSITE = 'WEBSITE',
}

// 图片来源
export enum ImageSource {
  LOCAL = 'local',
  URL = 'url',
  S3 = 's3',
  R2 = 'r2',
}

// 评论服务提供者
export enum CommentProvider {
  TWIKOO = 'twikoo',
  ARTALK = 'artalk',
  WALINE = 'waline',
  GISCUS = 'giscus',
}

// S3 Service Provider
export enum S3Provider {
  AWS = 'aws',
  ALIYUN = 'aliyun',
  TENCENT = 'tencent',
  MINIO = 'minio',
  OTHER = 'other', // 其它默认按照 MINIO 处理
}

// OAuth2 Provider
export enum OAuth2Provider {
  GITHUB = 'github',
  GOOGLE = 'google',
}

// Follow Status
export enum FollowStatus {
  NONE = 'none',
  PENDING = 'pending',
  ACCEPTED = 'accepted',
  REJECTED = 'rejected',
}

// Online Music Service Provider
export enum MusicProvider {
  NETEASE = 'netease', // 网易云音乐
  QQ = 'tencent', // QQ音乐
  APPLE = 'apple', // Apple Music
}

// Access Token Expiration Time
export enum AccessTokenExpiration {
  EIGHT_HOUR_EXPIRY = '8_hours', // 8小时
  ONE_MONTH_EXPIRY = '1_month', // 1个月
  NEVER_EXPIRY = 'never', // 永不过期
}
