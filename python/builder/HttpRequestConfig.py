from dataclasses import dataclass, field
from typing import Dict, Optional, Any, Literal

HttpMethod = Literal["GET", "POST", "PUT", "PATCH", "DELETE"]

@dataclass
class HttpRequestConfig:
    url: str
    method: HttpMethod = "GET"
    headers: Dict[str, str] = field(default_factory=dict)
    body: Optional[Any] = None
    timeout: int = 5  # segundos
    retries: int = 0
